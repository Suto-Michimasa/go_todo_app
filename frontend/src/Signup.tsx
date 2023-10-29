import React from 'react';
import { Box, Heading, Input, Button } from '@chakra-ui/react';
import Layout from './components/Layout';

const Signup: React.FC = () => {

  return (
    <Layout>
      <Heading as="h1" textAlign="center" mb="12">Signup</Heading>
      <Box as="form" method="post" action="/signup" className="form-signin">
        <Input mb="2" placeholder="お名前" name="name" required autoFocus />
        <Input mb="2" placeholder="Email" name="email" required type="email" />
        <Input mb="4" placeholder="パスワード" name="password" required type="password" />
        <Button type="submit" colorScheme="blue" w="100%">登録</Button>
      </Box>
    </Layout>
  );
};

export default Signup;
