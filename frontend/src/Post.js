import React, { Component } from 'react';


class CreateNew extends Component {
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
        <div>
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

export default CreateNew;