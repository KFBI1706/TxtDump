import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Post from './Post'
import Display from './Display'



function PageAction(props) {
  const action = props.action;
  if (action === "CreatePost"){
    return <Post />;
  }
  return <Display />
}
class App extends Component {
  
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          
        </header>
        <p>{this.props.testData}</p>
      <PageAction action="CreatePost" />
      
      </div>
    );
  }
}


export default App;