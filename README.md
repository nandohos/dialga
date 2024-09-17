# Dialga - A job scheduler in go

## Planned Features

[ ] **Job Queue**

- Data structure to store pending jobs

[ ] **Job Scheduling**

- Ability to schedule jobs for immediate or future execution

[ ] **Worker Process**

- Component that pulls jobs from the queue and executes them

[ ] **Item Queue**

- Ability to add items to the queue to be processed

[ ] **Job Storage**

- Method to persist job information, possibly using a database

[ ] **Basic Job Types**

- Support for fire-and-forget jobs (run once)
- Support for recurring jobs

[ ] **Error Handling**

- Mechanism to handle and log job execution failures

[ ] **Job Status Tracking**

- Track the state of jobs (queued, running, completed, failed)

[ ] **Simple API**

- For adding jobs to the queue and managing scheduled jobs
