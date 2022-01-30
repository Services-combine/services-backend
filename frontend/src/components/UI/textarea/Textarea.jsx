import React from 'react';
import classes from './Textarea.module.css'

const Textarea = (props) => {
    return (
      <textarea className={classes.Textarea} {...props} ></textarea>
    );
};

export default Textarea;