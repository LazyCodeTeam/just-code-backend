-- name: GetAllTechnologies :many
SELECT technology.*, 
section.id as section_id, 
section.position, 
section.title as section_title, 
section.description as section_description, 
section.image_url as section_image_url
FROM technology
LEFT JOIN section ON section.technology_id = technology.id
ORDER BY technology.position ASC, section.position ASC;
