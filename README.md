英文| [中文](README_ZH.md) 

# gDAG
gDAG is a platform for schedule and monitor workflows. gDAG is a task scheduling tool that manages task flows using Directed acyclic graph (DAG). Task scheduling can be achieved by setting task dependencies without knowing service data content.


## Directory structure
- control.sh Starts the control file
- app scheduler, worker and other app entries
- config Test configuration directory
- bin Indicates binary output
- docs Project description document
- lib Base libraries that do not depend on each other or the outer layer
- service Indicates the logical service layer

sh Function description
1. build an app. By default, all applications in the app directory are built without the build parameter. The specified applications are built with parameters and stored in the bin directory
> `#sh control.sh build dashboard`

2. Clean up all products in the bin
> `#sh control.sh clean`

3. Start the local doc file on port 8001
> `#sh control.sh doc`
> `http://localhost:8001/pkg/gDAG/`