import React from "react";
import ReactDOM from "react-dom";

import "./index.scss";
import Display from "./Display";

const App = () => (
    <Display/>
);
ReactDOM.render(<App/>, document.getElementById("app"));
