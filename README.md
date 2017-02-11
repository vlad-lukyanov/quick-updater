# quick-updater

quick-updater is simple step-by-step executing cmd that I use as updater for other programs. It uses json file as commands source

```json
{
 "steps":[{
   "command":"ls",
   "args":[],
   "skip_errors":true
  }, {
   "command":"/bin/bash",
   "args":["-c", "echo Hello from step 1!"]
  }]
} 
```
