class App extends React.Component {
    setState() {
        if (localStorage.getItem("access_token")) {
            this.loggedIn = true;
        } else {
            this.loggedIn = false;
        }
    }
    componentWillMount() {
        this.setState();
    }

    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {

    constructor(props) {
        super(props);
        this.authenticate = this.authenticate.bind(this);
    }

    authenticate() {
        console.log("AUTHENTICATE")
        localStorage.setItem("access_token", "YES NENA!");
        location.reload(true);

    }

    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Generative Cloud Services</h1>
                    <p>Sign in to get access </p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }

}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            workspaces: []
        }
    }

    serverRequest() {
        $.get("http://localhost:3000/api/workspaces", res => {
            this.setState({
                workspaces: res
            });
        });
    }

    componentDidMount() {
        this.serverRequest();
    }

    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>Workspaces</h2>
                    <p>Let's start code!!!</p>
                    <div className="row">
                        {this.state.workspaces.map(function(workspace, i){
                            return (<Workspace key={i} workspace={workspace} />);
                        })}
                    </div>
                </div>
            </div>
        )
    }
}

class Workspace extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            joined: ""
        }
        this.join = this.join.bind(this);
    }

    join() {
        console.log("gggg")
    }

    render() {
        return (
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading">#{this.props.workspace.id} <span className="pull-right">{this.state.joined}</span></div>
                    <div className="panel-body">
                        {this.props.workspace.info}
                    </div>
                    <div className="panel-footer">
                        {this.props.workspace.members} JoinED &nbsp;
                        <a onClick={this.join} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbs-up"></span>
                        </a>
                    </div>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));
