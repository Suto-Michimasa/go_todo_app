import React from 'react';
import { Box, Button, FormControl, FormLabel, Heading, Input, Center, Link } from '@chakra-ui/react';

const Login: React.FC = () => {
  return (
    <Center height="100vh">
      <Box as="form" action="/authenticate" method="post" width={['90%', '80%', '60%', '40%']} p={4} borderWidth={1} borderRadius={8} boxShadow="lg">
        <Heading as="h1" textAlign="center" marginBottom={4}>Login</Heading>

        <FormControl id="email" marginTop={4}>
          <FormLabel>Email</FormLabel>
          <Input type="email" name="email" isRequired placeholder="Email" />
        </FormControl>

        <FormControl id="password" marginTop={4}>
          <FormLabel>パスワード</FormLabel>
          <Input type="password" name="password" placeholder="パスワード" />
        </FormControl>

        <Button type="submit" width="full" colorScheme="blue" marginTop={4}>ログイン</Button>

        <Center marginTop={4}>
          <Link href="/signup">登録</Link>
        </Center>
      </Box>
    </Center>
  );
}

export default Login;
