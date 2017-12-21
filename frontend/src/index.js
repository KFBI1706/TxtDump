import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
var api = require('./RequestID.js')

api.getPost(11, "http://localhost:1337");
ReactDOM.render(<App />, document.getElementById('root'));

registerServiceWorker();
