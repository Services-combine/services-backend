import React from 'react'
import PostItem from './PostItem';
import { TransitionGroup, CSSTransition } from 'react-transition-group';

const PostList = ({posts, remove}) => {
    return (
        <div>
			<TransitionGroup>
				{posts.map(post => 
					<CSSTransition
						key={post.id}
						timeout={500}
						className="post"
					>
						<PostItem remove={remove} post={post} key={post.id} />
					</CSSTransition>
				)}
			</TransitionGroup>
		</div>
    );
};

export default PostList;
