import React from 'react';
import { Layout, Menu, Icon, Card, Row, Col, Breadcrumb } from 'antd';
import { Link, Route, Redirect } from 'react-router-dom'
import './dashboard.css';


import Product from '../product/product';
import Activity from '../activity/activity';
import Community from '../community/community';

const { SubMenu } = Menu;
const { Header, Sider, Content, Footer } = Layout;
// const { Link, HashRouter, Route, Switch, Redirect } = ReactRouterDOM;

class DashBoard extends React.Component {
  constructor(props) {
    super(props);

  }

  onCollapse() {
    this.setState({ collapsed: !this.state.collapsed });
  }

  render() {
    return (
      <div>
        <Layout>
          <Header className="header">
            <div className="logo-title" >
              XGoPkg
            </div>
            <Menu
              theme="dark"
              mode="horizontal"
              defaultSelectedKeys={['1']}
              style={{ lineHeight: '48px' }}
            >
              <Menu.Item key="1">
                <Link to='/dashboard/product'>产品</Link>
              </Menu.Item>
              <Menu.Item key="2">
                <Link to='/dashboard/activity'>活动</Link>
              </Menu.Item>
              <Menu.Item key="3">
                <Link to='/dashboard/community'>社区</Link>
              </Menu.Item>
            </Menu>
          </Header>

          <div>
            <Redirect exact from="/dashboard" to="/dashboard/product" />
            <Route exact path='/dashboard/product' component={Product} />
            <Route exact path='/dashboard/activity' component={Activity} />
            <Route exact path='/dashboard/community' component={Community} />
          </div>
        </Layout>
      </div>
    );
  }
}

export default DashBoard;