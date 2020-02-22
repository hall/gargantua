@Library('boxboat-dockhand@master')
import com.boxboat.jenkins.pipeline.build.*

properties([
  buildDiscarder(logRotator(numToKeepStr: '100'))
])

def build = new BoxBuild()

node() {
  build.wrap {
    gitlabCommitStatus {
      stage('Test'){
        build.composeBuild("sdk")
        sh '''
          docker run --rm -i hobbyfarm/gargantua-sdk:${GIT_COMMIT_SHORT_HASH:-dev} go test
        '''
      }
      stage('Build'){
        build.composeBuild("release")
      }
      stage ('Push'){
        build.push()
      }
    }
  }
}
