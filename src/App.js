import React from 'react';
import {Button, Card, Carousel, Icon, List, notification, PageHeader} from "antd";
import "antd/dist/antd.css";
import './App.css';

import config from './config'

var data = [];

class App extends React.Component {
    constructor(props) {
        super(props);
        this.setState({
            data: {}
        });
        this.request()
    }

    request = () => {
        fetch("/api/get-files", {
            method: "POST"
        })
            .then(res => res.json())
            .then(json => {
                data = json;
                this.setState({})
            })
            .catch(e => {
                console.error(e)
            })
    };

    render() {
        return (
            <div>
                <PageHeader onBack={() => null} backIcon={(<Icon type="menu"></Icon>)} title={config.appName}
                            subTitle={config.subtitle}>
                    <List
                        grid={{gutter: 16, column: 4}}
                        loading={false}
                        bordered={true}
                        dataSource={data}
                        renderItem={item => (
                            <Card title={item.title}>
                                <Carousel dotPosition={"bottom"} autoplay={true} autoplaySpeed={5000}>
                                    {
                                        item.files.map(v => {
                                            return (
                                                <div>
                                                    <Button type="primary" onClick={
                                                        () => {
                                                            notification.open({
                                                                message: config.notification_title,
                                                                description:
                                                                config.notification_message
                                                            });
                                                        }
                                                    } href={v.downloadURL}>{v.name}</Button>
                                                </div>)
                                        })
                                    }
                                </Carousel>
                            </Card>
                        )}
                    />
                </PageHeader>
            </div>
        );
    }
}

export default App;
