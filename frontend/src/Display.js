import React, { Component } from 'react';

class Display extends Component {
    render() {
      return(
        <div className="Content-body">
            <p>{this.props.body}</p>
        </div>
        
      );
    }
}

export default Display;