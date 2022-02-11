import React from 'react'
import AccountItem from './AccountItem';

const AccountList = ({accounts, remove, sendCodeParsing, parsingApi, sendCodeSession, createSession}) => {
    return (
        <div className='accounts container'>
            {accounts.map((account, index) => 
                <AccountItem remove={remove} account={account} sendCodeParsing={sendCodeParsing} parsingApi={parsingApi} sendCodeSession={sendCodeSession} createSession={createSession} index={index} key={index} />
            )}
        </div>
    );
};

export default AccountList;
