# 🦊 Identify, analyze, action! Deep monitoring with CI

- [Webcast Registration and Archive](https://page.gitlab.com/deep-monitoring-ci.html)
- [Slides](https://docs.google.com/presentation/d/1ONwIIzRB7GWX-WOSziIIv8fz1ngqv77HO1yVfRooOHM/edit?usp=sharing)

## Resources

- [GitLab CI Pipelines Exporter for Prometheus](https://github.com/mvisonneau/gitlab-ci-pipelines-exporter)
- [Performance Monitoring](https://docs.gitlab.com/ee/administration/monitoring/performance/)
- [GitLab.com Monitoring](https://about.gitlab.com/handbook/engineering/monitoring/)


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



