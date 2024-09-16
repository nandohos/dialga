# Dialga - A job scheduler in go

## Planned Features

1. **Job Queue**

   - Data structure to store pending jobs

2. **Job Scheduling**

   - Ability to schedule jobs for immediate or future execution

3. **Worker Process**

   - Component that pulls jobs from the queue and executes them

4. **Item Queue**

- Ability to add items to the queue to be processed

5. **Job Storage**

   - Method to persist job information, possibly using a database

6. **Basic Job Types**

   - Support for fire-and-forget jobs (run once)
   - Support for recurring jobs

7. **Error Handling**

   - Mechanism to handle and log job execution failures

8. **Job Status Tracking**

   - Track the state of jobs (queued, running, completed, failed)

9. **Simple API**
   - For adding jobs to the queue and managing scheduled jobs
