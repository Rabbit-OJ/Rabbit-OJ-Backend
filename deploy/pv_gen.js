const decl = (id) => `
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv${id}
  labels:
    type: local
spec:
  persistentVolumeReclaimPolicy: Delete
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /mnt/data${id}
`

const GENERATE_COUNT = 16;
let output = [];

for (let i = 1; i <= GENERATE_COUNT; i++) {
    output.push(decl(i));
}

const fs = require('fs');
fs.writeFileSync('./pv.yaml', output.join(`\n---\n`));
