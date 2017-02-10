--box.cfg{listen=3311}
box.schema.user.create('golang', {password='pass', if_not_exists = true})
s = box.schema.space.create('status', {if_not_exists=true})
s:create_index("command_id", {unique=true, parts={1, 'string'}})
s:create_index("user_id", {unique=false, parts={2, 'string'}})
box.schema.user.grant('golang', 'read,write,execute', 'universe')
fiber = require('fiber')

function update(commandId, userId, created, statusId, initiator)
	box.space.status:upsert({commandId, userId, fiber.time64(), statusId, initiator}, {{'=', 3, fiber.time64()}, {'=', 4, statusId}, {'=', 5, initiator}})
end

