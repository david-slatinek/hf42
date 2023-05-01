import React from "react";

import "./index.scss";
import {createRoot} from "react-dom/client";
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Detail from "./Detail";

const App = () => (
    <Router>
        <div>
            <Routes>
                <Route path="/:isbn" element={<Detail/>}/>
            </Routes>
        </div>
    </Router>
);

const rootElement = document.getElementById("app")!;
const root = createRoot(rootElement);
root.render(<App/>);
