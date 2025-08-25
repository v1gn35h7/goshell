pipeline {
    agent any // Specifies that the pipeline can run on any available agent
  tools { go '1.19' }

    stages {        
        stage('Build') {
            steps {
                sh 'go install github.com/jstemmer/go-junit-report/v2@latest'
                sh 'go mod tidy' // Ensures module dependencies are in sync
                sh 'go build -v ./...' // Builds your Go application
            }
        }

       stage('Test & Report') {
            steps {
                echo 'Running unit tests and generating JUnit report'
                // Run tests verbosely and pipe the output to go-junit-report
                sh 'go test -v ./... 2>&1 | go-junit-report -set-exit-code > test-report.xml' 
            }
            post {
                always {
                    // Publish JUnit test results to Jenkins
                    junit('test-report.xml') 
                }
            }
        }


    }


}
