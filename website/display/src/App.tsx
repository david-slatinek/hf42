import React from "react";

import "./index.scss";
import {createRoot} from "react-dom/client";
import Display from "./Display";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";

const App = () => (
    <Router>
        <div>
            <Routes>
                <Route path="/" element={<Display/>}/>
            </Routes>
        </div>
    </Router>
);

const rootElement = document.getElementById("app")!;
const root = createRoot(rootElement);
root.render(<App/>);
