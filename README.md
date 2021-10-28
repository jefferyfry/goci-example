# goci-example

[![CIS](https://app.soluble.cloud/api/v1/public/badges/56aaff94-9efc-4678-b863-3188f6c0226e.svg)](https://app.soluble.cloud/repos/details/github.com/jefferyfry/goci-example)  [![IaC](https://app.soluble.cloud/api/v1/public/badges/3c3fc247-bdc5-4382-b590-d96ca59a7efc.svg)](https://app.soluble.cloud/repos/details/github.com/jefferyfry/goci-example)  

This is a repo that is used for the blog [_Jenkins and JFrog Pipelines: Working Together to Release Your Software_](https://jfrog.com/blog/). The blog shows how to use a Jenkins Pipeline as CI and then use it to trigger JFrog Pipelines for staging, test and release. The application under tests is a simple GO REST application. The Jenkins pipeline is defined in the [Jenkinsfile](./Jenkinsfile) and the JFrog Pipeline is defined in the [pipeline.yml](./pipeline.yml).

![Diagram-all](https://user-images.githubusercontent.com/6440106/80759359-09d39e80-8aec-11ea-9beb-0334a81b2e50.png)

## Jenkins

![Jenkins](https://user-images.githubusercontent.com/6440106/80759523-4c957680-8aec-11ea-84d5-f29a9828cd1e.png)

## JFrog Pipelines

![JFrog Pipelines](https://user-images.githubusercontent.com/6440106/80759614-6b940880-8aec-11ea-9a9b-d1447b4dad81.png)
