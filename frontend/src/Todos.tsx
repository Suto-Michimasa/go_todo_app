import React from "react";
import { Box, Button, Heading } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import Layout from './components/Layout';

const Todos: React.FC = () => {
  const navigate = useNavigate();

  return (
    <Layout>
      <Box>
        <Heading mb={6}>Index</Heading>
      </Box>
    </Layout>
  );
};

export default Todos;
