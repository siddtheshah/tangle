/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * Generated with the TypeScript template
 * https://github.com/react-native-community/react-native-template-typescript
 *
 * @format
 */

import React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {SafeAreaView, StyleSheet, Text, Button, View} from 'react-native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import HomeScreen from './home';
import ProfileScreen from './profile';
import BottomNav from './bottom';

const Stack = createNativeStackNavigator();

const App = () => {
  return (
    <SafeAreaView style={styles.background}>
      <NavigationContainer>
        <Stack.Navigator initialRouteName="Bottom">
          <Stack.Screen
            name="Bottom"
            component={BottomNav}
            options={{title: 'Bottom'}}
          />
          <Stack.Screen
            name="Home"
            component={HomeScreen}
            options={{title: 'Welcome'}}
          />
          <Stack.Screen
            name="Details"
            component={ProfileScreen}
            options={{title: 'Details'}}
          />
        </Stack.Navigator>
      </NavigationContainer>
    </SafeAreaView>
  );
};

export default App;

const styles = StyleSheet.create({
  background: {
    flex: 1,
    backgroundColor: 'white',
    fontSize: '30px',
  },
});
