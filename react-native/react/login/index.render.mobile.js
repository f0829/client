import React, {Component} from 'react'
import {StyleSheet, View} from 'react-native'
import ProgressIndicator from '../common-adapters/progress-indicator'

export default class LoginRender extends Component {
  render () {
    return (
      <View style={styles.container}>
        <ProgressIndicator/>
      </View>
    )
  }
}

LoginRender.propTypes = {
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center'
  }
})
