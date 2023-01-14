pipeline {
  options {
    buildDiscarder(logRotator(numToKeepStr: "15", artifactNumToKeepStr: "15"))
    disableConcurrentBuilds()
  }
  agent any
  stages {
    stage("scm") {
      steps {
        git(url: "git@nuc.lliu.ca:app/weather_app_fiber.git", branch: "master", poll: true)
      }
    }

    stage("test") {
      steps {
        sh 'pwd'
        sh 'ls -al'
        //sh 'make test'
      }
    }
    // stage("coverage") {
    //   steps {
    //     sh 'sonar-scanner -Dproject.settings=./sonar-project.properties'
    //   }
    // }
    // stage("package") {
    //   steps {
    //     sh 'make rpm'
    //   }
    // }
    stage("Build Docker Image") {
      steps {
        sh "docker build -t nuc.lliu.ca/lliu/weather_app_fiber:${BUILD_NUMBER} ."
      } 
    }
    stage("Docker Login and Push") {
      steps{
        withCredentials([string(credentialsId: "dockerHubPass", variable: "dockerHubPass")]) {
          sh "docker login nuc.lliu.ca -u lliu -p $dockerHubPass"
        }
        sh "docker push nuc.lliu.ca/lliu/weather_app_fiber:${BUILD_NUMBER}"
      }
    }
    stage('Build Completed') {
      steps {
        sh '''
        echo "Build has completed..."
        '''
      }
    }
  }
}