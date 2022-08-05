import React from 'react';
import {Text, Button, View} from 'react-native';
import {StyleSheet} from 'react-native';
import BottomNav from './bottom';

export default function HomeScreen({navigation}) {
  return (
    <>
      <View style={styles.navigator}>
        <Text>Home Screen</Text>
      </View>
      <BottomNav navigation={navigation} />
    </>
  );
}
const styles = StyleSheet.create({
  navigator: {
    flex: 1,
    textAlign: 'center',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
