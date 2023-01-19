import { Button, Checkbox, Form, Input } from 'antd';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import { useCookies } from 'react-cookie'



const headers = {
  'Content-Type': 'text/plain'
};

const LogIn = () => {
  
  const navigate = useNavigate();

  const [cookies, setCookie] = useCookies(['token'])
  
  const onFinish = (values) => {
    console.log('Success:', values);

    axios.post('http://localhost:8080/v1/users/login', {
      email: values.username,
      password: values.password,
    },
    {headers}
    )
    .then(function (response) {
      console.log(response);
      setCookie('token', response.data.token, { path: '/' })
      navigate("/");
    })
    .catch(function (error) {
      console.log(error);
    });
  };
  
  const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };

  
  return (
    <Form
      name="basic"
      labelCol={{
        span: 8,
      }}
      wrapperCol={{
        span: 16,
      }}
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[
          {
            required: true,
            message: 'Please input your username!',
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[
          {
            required: true,
            message: 'Please input your password!',
          },
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        name="remember"
        valuePropName="checked"
        wrapperCol={{
          offset: 8,
          span: 16,
        }}
      >
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item
        wrapperCol={{
          offset: 8,
          span: 16,
        }}
      >
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
};

export default LogIn;
