# Dialga - A job scheduler in go

## Planned Features

1. **Job Queue**

   - Data structure to store pending jobs

2. **Job Scheduling**

   - Ability to schedule jobs for immediate or future execution

3. **Worker Process**

   - Component that pulls jobs from the queue and executes them

4. **Job Storage**

   - Method to persist job information, possibly using a database

5. **Basic Job Types**

   - Support for fire-and-forget jobs (run once)
   - Support for recurring jobs

6. **Error Handling**

   - Mechanism to handle and log job execution failures

7. **Job Status Tracking**

   - Track the state of jobs (queued, running, completed, failed)

8. **Simple API**
   - For adding jobs to the queue and managing scheduled jobs
