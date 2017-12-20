import * as React from 'react';
import './App.css';

const logo = require('./logo.svg');

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </div>
        <p className="App-intro">
          Fill in ur paste mate:
        </p>
        <form>
          <input className="textField" type="text"></input>
          <br />
          <button>Submit</button>
        </form>
      </div>
    );
  }
}

export default App;
