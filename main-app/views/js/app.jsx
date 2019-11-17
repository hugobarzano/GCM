const AUTH0_CLIENT_ID = "egdgErlFN4eEeUmWX3an6jKjx1pNpmNT";
const AUTH0_DOMAIN='cesarcorp.eu.auth0.com';
const AUTH0_CALLBACK_URL = location.href;
const AUTH0_API_AUDIENCE = "https://cesarcorp.eu.auth0.com/api/v2/";

class App extends React.Component {
    parseHash() {
        this.auth0 = new auth0.WebAuth({
            domain: AUTH0_DOMAIN,
            clientID: AUTH0_CLIENT_ID
        });
        this.auth0.parseHash(window.location.hash, (err, authResult) => {
            if (err) {
                return console.log(err);
            }
            if (
                authResult !== null &&
                authResult.accessToken !== null &&
                authResult.idToken !== null
            ) {
                localStorage.setItem("access_token", authResult.accessToken);
                localStorage.setItem("id_token", authResult.idToken);
                localStorage.setItem(
                    "profile",
                    JSON.stringify(authResult.idTokenPayload)
                );
                window.location = window.location.href.substr(
                    0,
                    window.location.href.indexOf("#")
                );
            }
        });
    }

    setup() {
        $.ajaxSetup({
            beforeSend: (r) => {
                if (localStorage.getItem("access_token")) {
                    r.setRequestHeader(
                        "Authorization",
                        "Bearer " + localStorage.getItem("access_token")
                    );
                }
            }
        });
    }

    setState() {
        let idToken = localStorage.getItem("id_token");
        if (idToken) {
            this.loggedIn = true;
        } else {
            this.loggedIn = false;
        }
    }

    componentWillMount() {
        this.setup();
        this.parseHash();
        this.setState();
    }

    render() {
        if (this.loggedIn) {
            return <LoggedIn />;
        }
        return <Home />;
    }
}

class Home extends React.Component {

    constructor(props) {
        super(props);
        this.authenticate = this.authenticate.bind(this);
    }

    authenticate() {
        this.WebAuth = new auth0.WebAuth({
            domain: AUTH0_DOMAIN,
            clientID: AUTH0_CLIENT_ID,
            scope: "openid profile",
            audience: AUTH0_API_AUDIENCE,
            responseType: "token id_token",
            redirectUri: AUTH0_CALLBACK_URL
        });
        this.WebAuth.authorize();
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
        this.serverRequest = this.serverRequest.bind(this);
        this.logout = this.logout.bind(this);
    }

    logout() {
        localStorage.removeItem("id_token");
        localStorage.removeItem("access_token");
        localStorage.removeItem("profile");
        location.reload();
    }

    serverRequest() {
        $.get("http://localhost:3000/back/workspaces", res => {
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
