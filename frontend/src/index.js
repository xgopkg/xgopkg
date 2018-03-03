import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter ,Router, Route, Link,hasHistory,Switch } from 'react-router-dom'
import './index.css';
// import App from './App';
import registerServiceWorker from './registerServiceWorker';
import DashBoard from './components/dashboard/dashboard'


const ROUTE_CONFIG = () => (
  <main>
      <Switch>
          <Route path="/" exact component={DashBoard} />
          <Route path="/dashboard" component={DashBoard} />
      </Switch>
  </main>
)


ReactDOM.render((
  <HashRouter>
    <ROUTE_CONFIG />
  </HashRouter>
), document.getElementById('root'));
registerServiceWorker();
