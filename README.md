# quick-updater

quick-updater is simple step-by-step executing cmd that I use as updater for other programs. It uses json file as commands source
also it can replace "mustashed variables" from config with arguments from command line

## run
```sh
./quick-updater arg0 arg1 "long arg2"
```
## config
```json
{
"steps":[
 {
  "command":"ls",
  "args":[],
  "skip_errors":true
 },
 {
  "command":"/bin/bash",
  "args":["-c", "echo {{0}} from step 1!"]
 },
 {
  "command":"/bin/bash",
  "args":["-c", "echo {{1}} from step 2!"]
 },
 {
  "command":"/bin/bash",
  "args":["-c", "echo {{2}} from step 3!"]
 },
 {
  "command":"/bin/bash",
  "args":["-c", "echo {{3}} from step 4!"]
 }
]
} 

```
