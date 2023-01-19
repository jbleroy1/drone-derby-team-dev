# drone-derby

Monorepo for the Drone Derby Project.

For more information see [go/cloud-drone-derby](http://go/cloud-drone-derby).

This repo has the following directory structure:

* [internal/](internal/): Use this directory for all internal code that **will not** be shared with hackathon attendees.
* [external/](external/): Use this directory for all code that **will be** shared with hackathon attendees.

There should be no commits outside of these two directories.

## Development Guiding Principles

Some of our biggest customers and some of the public will end up seeing the code, environment and infrastructure we create. Some of those that attend the hackathon are likely to be technical and some technical leads for our customers. It is important that we hold what we do to a high standard, that expected of engineering from Google.

It’s ok to be fast and scrappy to test something out, but the code we ship for the hackathon should:
* Include documentation, both inline in the code and a sensible readme for each component
* Should lint to a Google Coding standard
* Include unit and integration tests (where appropriate) including end 2 end tests of the frontend critical user journeys.
* Implement cloud best practices, such as infrastructure as code
* Each deployable artefact should have a Cloud Build or Cloud Deploy pipeline showing best practice software delivery in Google Cloud
* We will review each other’s work, be open, honest and most of all Googley in the approach
* Ensure all files have a blank new line at the end
