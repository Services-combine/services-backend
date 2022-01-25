import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import PostService from '../API/PostService';
import Loader from '../components/UI/loader/Loader';
import { useFetching } from '../hooks/useFetching';

const PostIdPage = () => {
    const params = useParams();
    const [post, setPost] = useState({});
    const [comments, setComments] = useState([]);

    const [fetchPostById, isLoading, error] = useFetching(async () => {
        const response = await PostService.getById(params.id)
        setPost(response.data);
    })
    const [fetchComments, isComLoading, comError] = useFetching(async () => {
        const response = await PostService.getCommentByPostId(params.id)
        setComments(response.data);
    })

    useEffect(() => {
        fetchPostById()
        fetchComments()
    }, [])

    return (
        <div>
            <h2>Описание поста ID = {params.id}</h2>
            {isLoading
                ? <Loader/>
                : <div>{post.id}. {post.title}</div>
            }

            <h2 style={{marginTop: 50}}>Комментарии</h2>
            {isComLoading
                ? <Loader/>
                : <div>
                    {comments.map(comm => 
                        <div key={comm.id} style={{marginTop: 15}}>
                            <h4>{comm.email}</h4>
                            <div>{comm.body}</div>
                        </div>
                    )}
                </div>
            }

        </div>
    );
};

export default PostIdPage;
