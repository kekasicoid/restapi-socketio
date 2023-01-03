import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1s', target: 10 },
    
  ],
};

export default function () { 
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };
  const url = 'http://localhost:8989/ping'
  const res = http.get(url, "", params);
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}