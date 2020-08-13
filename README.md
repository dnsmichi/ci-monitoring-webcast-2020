# 🦊 Identify, analyze, action! Deep monitoring with CI

- Webcast Registration and Archive - TODO
- Slides - TODO

## Features

### Test Reports on GitLab Pages

https://dnsmichi.gitlab.io/ci-monitoring-webcast-2020/

This requires:

- Pages enabled
- Tests rendered as HTML output
- Tests for MRs with a defined name
- An index.html tree builder for the Pages index site

```
unit-test:
  stage: test
  script:
    - go test -coverprofile=.coverage.cov $(go list ./... | grep -v /vendor/) 
  artifacts:
    paths:
    - .coverage.cov

test-coverage-report:
  stage: coverage
  needs: ["unit-test"]
  coverage: /regular total:\s+\(statements\)\s+\d+.\d+\%/
  variables:
    COV_FILENAME: coverage-$CI_COMMIT_BRANCH-job-$CI_JOB_ID.html
  script:
    - go tool cover -func .coverage.cov
    - go tool cover -html=.coverage.cov -o public/$(echo $COV_FILENAME | sed -e 's/\//___/g') # replace slashes in branch names
  artifacts:
    paths:
    - public/coverage*.html    
```



