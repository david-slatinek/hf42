import React from "react";
import "./index.scss";
import ReactDOM from "react-dom";

const App = () => (
    <div className="mt-10 text-3xl mx-auto max-w-6xl">
        <div>Name: detail</div>
        <div>Framework: react</div>
        <div>Language: TypeScript</div>
        <div>CSS: Tailwind</div>
    </div>
);
ReactDOM.render(<App/>, document.getElementById("app"));
