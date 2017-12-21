import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import {getPost} from './RequestID'

getPost(11, "http://localhost:1337", function(content){
    console.log(content)
});

ReactDOM.render(<App />, document.getElementById('root'));

registerServiceWorker();
