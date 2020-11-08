
-- configure tarantool
box.cfg{
    listen = 3301
}

box.once('init', function()
    s = box.schema.space.create('sessions')
    s:format({
        {name = 'session_value', type = 'string'},
        {name = 'user_id', type = 'unsigned'}
    })

    test_sessions = box.schema.space.create('test_sessions')
    test_sessions:format({
        {name = 'session_value', type = 'string'},
        {name = 'user_id', type = 'unsigned'}
    })

    s:create_index('primary', {type = 'HASH', parts = {'session_value'}})

    box.schema.func.create('create_session', {setuid= true})
    box.schema.func.create('check_session', {setuid= true})
    --test functions
    box.schema.func.create('create_test_session', {setuid= true})
    box.schema.func.create('check_test_session', {setuid= true})
    --main user
    box.schema.user.create("Backend_cinema_interface", {password='some_password'})
    box.schemwa.user.grant('Backend_cinema_interface', 'read,write,execute,create,drop', 'universe')
    box.schema.user.grant('Backend_cinema_interface', 'execute', 'function', 'create_session')
    box.schema.user.grant('Backend_cinema_interface', 'execute', 'function', 'check_session')
    box.session.su('Backend_cinema_interface')

    --test user

    box.schema.user.create("Backend_cinema_test_user", {password='some_password'})
    box.schemwa.user.grant('Backend_cinema_test_user', 'read,write,execute,create,drop', 'universe')
    box.schema.user.grant('Backend_cinema_test_user', 'execute', 'function', 'create_test_session')
    box.schema.user.grant('Backend_cinema_test_user', 'execute', 'function', 'check_test_session')
    box.test_sessions.su('Backend_cinema_test_user')

    print("tarantool initialized")
end)


function create_session(cookie_value, cookie_info, user_id)
    print('received data', user_id, cookie_info)
    box.space.sessions:insert{cookie_value, user_id, cookie_info}

    return cookie_value
end

function check_session(cookie_value)
    local query_result = box.space.sessions:select{cookie_value}[1]
    print('session found', query_result)

    return query_result
end

function create_test_session(cookie_value, cookie_info, user_id)
    print('received data', user_id, cookie_info)
    box.space.test_sessions:insert{cookie_value, user_id, cookie_info}

    return cookie_value
end

function check_test_session(cookie_value)
    local result = box.space.test_sessions:selec{cookie_value}[1]
    return result
end