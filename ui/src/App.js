import React from "react";
import Layout from "./components/Layout/Layout";
import Navigationbar from "./components/Navigation/NavigationBar";
import Login from "./components/Login/Login";
import UserList from "./components/User/UserList";
import Home from "./components/Home";
import FileList from "./components/File/FileList";
import FileListFunctional from "./components/File/FileListFunction";
import NoMatch from "./components/NoMatch";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { AuthProvider } from "./hooks/useAuthContext";
import { PrivateRoute } from "./components/PrivateRoute";
import "bootstrap/dist/css/bootstrap.min.css";

export default function App() {
  return (
    <AuthProvider>
      <Router>
        <Navigationbar />
        <Layout>
          <Switch>
            <PrivateRoute exact path="/" component={Home} />
            <PrivateRoute path="/fileList" component={FileList} />
            <PrivateRoute
              path="/fileListFunctional"
              component={FileListFunctional}
            />
            <PrivateRoute path="/userList" component={UserList} />
            <Route path="/login" component={Login} />
            <Route component={NoMatch} />
          </Switch>
        </Layout>
      </Router>
    </AuthProvider>
  );
}
