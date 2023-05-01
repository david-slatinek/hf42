import React from "react";

import "./index.scss";
import {createRoot} from "react-dom/client";
import Display from "./Display";

const App = () => (
    <div>
        <Display/>
    </div>
);

const rootElement = document.getElementById("app")!;
const root = createRoot(rootElement);
root.render(<App/>);
