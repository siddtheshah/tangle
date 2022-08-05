import { CurrentRenderContext } from '@react-navigation/native';
import React from 'react';
import {Text, Button, View} from 'react-native';
import {StyleSheet} from 'react-native';

export default function BottomNav({navigation}) {
  return (
    <View style={styles.bottom}>
      <Button
        title="Go to Detail"
        onPress={() => navigation.navigate('Details')}
      />
      <Button title="Go to Home" onPress={() => navigation.navigate('Home')} />
    </View>
  );
}

const styles = StyleSheet.create({
  bottom: {
    flex: 1,
    flexDirection: 'row',
    display: 'flex',
    textAlign: 'center',
    alignItems: 'flex-end',
    justifyContent: 'space-around',
  },
});
