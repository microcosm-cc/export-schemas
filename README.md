Forum Export Schemas
====================

A set of JSON schemas that will hopefully evolve into a common format to which forums from any platform (Discourse, Invision, Microcosm, moot.it, phpBB, Vanilla, vBulletin, XenForo, etc) can be exported.

The idea behind this is that any forum should be able to export core data into a common (non-destination specific) format. And that any other forum could write a single importer to import that data.

*The goal is data portability of forums.*

We're not going to get this right straight away, but over time and after implementing exports from multiple systems we'll get it more right than wrong.

We'll start with a scratch version of basic data, will evolve it towards a fuller schema.

To anyone wishing to help there are several specific goals with these schemas:

1. Consistency.
2. Simplicity of export and import.

For consistency we are following schema.org where we can, and only adding our own fields where necessary to extend it. Schema.org have already done a lot of the hard work of thinking about the structure of data that comprises a lot of the core of a forum (comments), we're standing on their shoulders.

Simplicity is achieved by going for a low-coupled/low-dependency export format. In essence this means that exporting a comment is limited to that, we don't go and export the user and any attachments at the same time... those will be different export files and it's up to the importer to do the job of determining the order to import something, the export exports as simple a structure as it can with few internal dependencies.

This repo will also provide Go structs to allow Go programmers to import the latest versions of the schemas to ensure that they are exporting and importing to the latest version.