import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    console.log("Text submitted:", this.state.value)
    event.preventDefault();
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          
        </header>
        <p className="App-intro">
          insert ur text kiddo
        </p>
      <form onSubmit={this.handleSubmit}>
        <input type="text" className="textField" value={this.state.value} onChange={this.handleChange} /> <br />
        <input type="submit" value="Submit" />
      </form>
      </div>
    );
  }
}

export default App;
