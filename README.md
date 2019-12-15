# Rabbit OJ’s Technical Details

## 🚗Technical Stacks
- Frontend: Angular & Rxjs & Typescript 
- Backend: Go & RabbitMQ & Docker & Web Socket & MySQL & gRPC & Protobuf

We will discuss the actions after the user submit a piece of code.

## 🤩System Design

- Frontend(Angular) <==> Server(GIN) <==> Message Queue(Rabbit MQ) <==> Judger

- Judger —> (Many) Machines —> Scheduler —> Compiler —> Runner —> Callback

## 🚄Work Flows

1. Frontend POST a file to the Server with its question id & language type & token, then the server will validate the request and decide if the language is supported.
2. The server will return a unique submission id to the client and the client will establish a web socket connection to the server
3. The server will serialize the message to the protobuf bytes and send it to the Message Queue Exchange
4. The Message Queue Exchange will match its route and enqueue to the “judge” Queue
5. An idle Consumer (Judger) will consume the message and decide if the language is supported.
    - If the language is not configured in the config file, nack it and re-queue the message to the queue
    - This case may happen when a new language is added to the server and the judgers are executing the rolling update action. The message is sent to an judger with old version.
    - Re-queue action will guarantee the message will be consumed correctly
6. The judger will check if it have the right version of test case, if not, dial to server and require the latest test case with credential
    - We tag each set of test cases with “version”, if the administrator update the test cases set, it will guarantee that the judger will always use the latest test case
7. If the language requires the code to compile first (for example: C / C++ / Java), a compiler container will be started and try to compile the code. If an error occurred, then the submission is considered to be CE
    - If the language doesn’t require a compile procedure (like python, node.js), this step will be jumped
8. Then we will start a special container built by ourselves called “tester”, It will run the code or binary file with special arguments
    - 🤔Time Limit: We will start a go-routine, the routine will kill the process if time limit the restriction and judge TLE
    - 🤔Space Limit: We will start a go-routine, parse the file content ‘/proc/{pid}/stat’ every 100ms , If it exceed the restriction, kill the process and judge MLE
    - 🤔If the process exit unexpectedly or doesn’t return zero, the submission will be considered to be RE
9. The test inputs are mounted to the tester container ,however, the right outputs are not. The judge process will be executed after the tester container exit. We have 3 modes to test if an output is AC or WA
    - Text Compare: Simply compare if the right output and the code’s output are equal
    - Stdout Compare: Split the output with separate chars like ‘\n’, ‘\r’, ‘ ’, ‘\t’. Then compare two arrays
    - Float Compare: this mode is similar with mode2. The only difference is that if abs(rightArr[i] - outputArr[i]) <= 1e-6 , the answer will be accepted.
10. The judger will serial a result protobuf object and send it to the Message Queue
11. The Server will consume the message, storage the status into the database, notify the client with the web socket
12. The client received the result and re-render the page

## 🤔How to guarantee that a message will always be consumed?
- If a nature disaster is happened, some messages will be lost.😖
- If the judger is killed with #9 interrupt (FORCE EXIT), The running submission’s result will lose. We should avoid to stop the judger with this signal.😖
- If the judger received #15 interrupt (Maybe Control-C ?), The idle machine go-routines and consumer will stop immediately (This will guarantee that no more messages will be consumed), the process will exit until all the running go-routines finish.☕️

## 🤔Should we be optimistic or pessimistic?
- If we are optimistic, some “unnecessary” check producers will be ignored, the performance will be higher, but maybe we will meet some unexpected errors in some special & extreme situations.
- If we are pessimistic, some “unnecessary” check producers will be performed, the performance will be lower, but when we face some special & extreme situations, we can handle these errors confidently.

- For example:
    - Check if required test cases are valid is performed in the scheduler, should we check them again in the tester?
    - When we received #15 interrupt signal from the operating system, should we be optimistic and consider that all the running processes will finish in a short time? Or we can be pessimistic and consider that the running processes will take a long long time, we re-queue the messages to the message queue and abort all the running processes?

## 👻A new language can be supported by only updating the config file!