import logo from './logo.svg';
import './App.css';
import { Button } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import InfoIcon from '@material-ui/icons/Info';
import axios from 'axios';
import { useState, useEffect } from 'react';


const useStyles = makeStyles((theme) => ({
  margin: {
    margin: theme.spacing(1),
  }
}));



function App() {
  let [value, setValue] = useState(0);
  useEffect(() => {
    document.title = `Counter: ${value}`;
  }, [value]);

  function incrementRequest() {
      axios.post("http://localhost:8000/increment", {
      value: value
    })
    .then(function (response) {
      setValue(response.data.value)
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });
  }

  function decrementRequest() {
      axios.post("http://localhost:8000/decrement", {
      value: value
    })
    .then(function (response) {
      setValue(response.data.value)
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });
  }

  function readValueRequest() {
      axios.post("http://localhost:8000/value", {
      value: value
    })
    .then(function (response) {
      setValue(response.data.value)
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });
  }

  function resetRequest() {
      axios.get("http://localhost:8000/reset", {
      value: value
    })
    .then(function (response) {
      setValue(response.data.value)
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });
  } 

  const styles = useStyles();

  return (
    <div className="App">
      <div className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          A simple counter service. Current counter: {value}
        </p>
        <span style={{ display: "inline-flex", alignItems: "center" }}>
          <InfoIcon></InfoIcon>
          <p style={{fontSize: 20, marginLeft: 10}}>
            Counter-value has a min/max range of [-10001, 10001].
          </p>
        </span>
        <div>
          <Button className={styles.margin} onClick={incrementRequest} variant="contained" color="primary">Increment +1</Button>
          <Button className={styles.margin} onClick={decrementRequest} variant="contained" color="primary">Decrement -1</Button>
          <Button className={styles.margin} onClick={resetRequest} variant="contained" color="primary">Reset counter to 0</Button>
        </div>
      </div>
    </div>
  );
}

export default App;
