import React from 'react';
import { Box, Heading, Input, Button, useColorModeValue } from '@chakra-ui/react';
import Layout from './components/Layout';

const Signup: React.FC = () => {
  const bgColor = useColorModeValue("#ccffcc", "#4A5568");

  return (
    <Layout>
      <Heading as="h1" textAlign="center" mb="4">Signup</Heading>
      <Box as="form" method="post" action="/signup" className="form-signin">
        <Heading as="h2" fontSize="xl" mb="2">
          SampleApp
        </Heading>
        <Box as="p" fontSize="lg" mb="4">登録</Box>
        <Input mb="2" placeholder="お名前" name="name" required autoFocus />
        <Input mb="2" placeholder="Email" name="email" required type="email" bg={bgColor} />
        <Input mb="4" placeholder="パスワード" name="password" required type="password" />
        <Button type="submit" colorScheme="blue" w="100%">登録</Button>
      </Box>
    </Layout>
  );
};

export default Signup;
