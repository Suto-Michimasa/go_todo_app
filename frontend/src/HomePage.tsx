import React from "react";
import { Box, Button, Center, Heading } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import Layout from './components/Layout';

const HomePage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <Layout>
      <Box>
        <Heading mb={6}>Go Website</Heading>
        <Button onClick={() => navigate("/signup")}>Go to Sign Up</Button>
      </Box>
    </Layout>
  );
};

export default HomePage;
