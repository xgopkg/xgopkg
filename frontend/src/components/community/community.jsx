import React from 'react';
import { Layout, Menu, Icon, Card, Row, Col, Breadcrumb } from 'antd';

const { SubMenu } = Menu;
const { Header, Sider, Content, Footer } = Layout;

class Community extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      collapsed: false,
    };
  }
  onCollapse() {
    this.setState({ collapsed: !this.state.collapsed });
  }
  render() {
    return (
      <div>
        <Layout>
          <Sider
            collapsible
            collapsed={this.state.collapsed}
            onCollapse={this.onCollapse.bind(this)}
          >
            <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
              <SubMenu
                key="sub1"
                title={<span><Icon type="user" /><span>社区一</span></span>}
              >
                <Menu.Item key="3">Tom</Menu.Item>
                <Menu.Item key="4">Bill</Menu.Item>
                <Menu.Item key="5">Alex</Menu.Item>
              </SubMenu>
              <SubMenu
                key="sub2"
                title={<span><Icon type="team" /><span>社区二</span></span>}
              >
                <Menu.Item key="6">Team 1</Menu.Item>
                <Menu.Item key="8">Team 2</Menu.Item>
              </SubMenu>
            </Menu>
          </Sider>
          <Layout style={{ padding: '0 24px 24px' }}>
            <Content style={{ margin: '0 16px' }}>
              <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>User</Breadcrumb.Item>
                <Breadcrumb.Item>Bill</Breadcrumb.Item>
              </Breadcrumb>
              <div style={{ padding: 24, background: '#fff', minHeight: '80vh' }}>
                <h1>社区页面</h1>
              </div>
            </Content>
          </Layout>
        </Layout>
      </div>
    )
  }
}

export default Community;


