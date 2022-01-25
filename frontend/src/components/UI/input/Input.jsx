import React from 'react';
import classes from './Input.module.css'

const Input = (props) => {
    return (
      <input className={classes.Input} {...props}/>
    );
};

export default Input;