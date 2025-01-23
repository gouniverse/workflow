# Workflow <a href="https://gitpod.io/#https://github.com/gouniverse/workflow" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/workflow/workflows/tests/badge.svg)

WorkFlow is a robust, asynchronous durable workflow task queue package designed to offload time-consuming or resource-intensive operations from your main application.

By deferring workflow tasks to the background, you can improve application responsiveness and prevent performance bottlenecks.

WorkFlow leverages a durable database (SQLite, MySQL, or PostgreSQL) to ensure reliable persistence and fault tolerance.

## License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0). You can find a copy of the license at [https://www.gnu.org/licenses/agpl-3.0.en.html](https://www.gnu.org/licenses/agpl-3.0.txt)

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## What is a workflow?

A workflow is a set of steps that when executed in specific order
complete a business process.

For instance:

- Create a user account
- Send a confirmation email
- Send a welcome email

## How it works

- Investigate what the business process entails.
- Define what the steps are and the order in which they should be executed,
  ensuring that each step is actionable and measurable.
- Create step definitions that outline the specific tasks, inputs, and outputs
  for each step.
- Create a workflow definition that outlines the steps and their relationships,
  using the step definitions created in step 3.
- To run the workflow, create a new workflow instance, based on the workflow 
  definition you already specified.
- Start the workflow instance, once any necessary conditions or triggers have been met.
- Monitor the workflow instance and wait for all steps to complete,
  checking the status of each step as it finishes to determine
  if it succeeded or failed.
- Review the outcome of the workflow instance and verify that it has completed
  successfully, taking into account the status of each individual step.
- If a step fails, inspect the logs to identify the cause of the failure,
  remove the obstacle or correct the issue.
- You should be able to rerun the failing step, and continue execution of
  the workflow instance from the point where the failure occurred,
  ensuring that any subsequent steps are executed in the correct order.



## Similar
============
- https://github.com/iTrellis/workflow/tree/main
- https://github.com/matryer/flower/blob/master/flower.go
- https://github.com/hibiken/asynq
- https://github.com/gocelery/gocelery
