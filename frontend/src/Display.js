import React, { Component } from 'react';

const displayAPI = "http://localhost:1337/post/"

class Display extends Component {
  constructor(props) {
    super(props);
  
    this.state = {
      content: [],
    };
  }

  componentDidMount() {
    fetch(displayAPI + "request/4838680")
      .then(response => response.json())
      .then(data => this.setState({content: data.Content}))
  }
    render() {
      return(
        <div className="Content-body">
            <p>{this.props.body}</p>
        </div>
        
      );
    }
}

export default Display;