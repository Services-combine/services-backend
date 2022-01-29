import React from 'react'

const CountAccounts = (props) => {
    return (
        <div className="count-accounts">
            <h6><i className="fas fa-user-alt"></i> - {props.all}</h6>
            <h6><i className="fas fa-check"></i> - {props.clean}</h6>
            <h6><i className="fas fa-info-circle"></i> - {props.block}</h6>
        </div>
    );
};

export default CountAccounts;
