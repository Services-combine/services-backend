import React from 'react';
import classes from './Error.module.css'

const Error = ({children, ...props}) => {
    return (
        <div {...props} className={classes.Err}>
            {children}
        </div>
    );
};

export default Error;