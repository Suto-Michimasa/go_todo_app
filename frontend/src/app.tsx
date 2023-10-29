import React from "react";
import { ChakraProvider } from "@chakra-ui/react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./HomePage";
import Signup from "./Signup";
import Login from "./Login";
import Todos from "./Todos";

function App() {
  return (
    <ChakraProvider>
      <Router>
        <Routes>
          <Route path="/signup" element={<Signup />} />
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<HomePage />} />
          <Route path="/todos" element={<Todos />} />
        </Routes>
      </Router>
    </ChakraProvider>
  );
}

export default App;
