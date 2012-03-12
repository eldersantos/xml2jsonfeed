#ifndef __CHELPER_H__
#define __CHELPER_H__

#include <libxml/tree.h>
#include <libxml/parser.h>
#include <libxml/HTMLtree.h>
#include <libxml/HTMLparser.h>
#include <libxml/xmlsave.h>
#include <libxml/xpath.h>

xmlDoc* xmlParse(void *buffer, int buffer_len, void *url, void *encoding, int options, void *error_buffer, int errror_buffer_len);
xmlNode* xmlParseFragment(void* doc, void *buffer, int buffer_len, void *url, int options, void *error_buffer, int error_buffer_len);
int xmlSaveNode(void *buffer, int buffer_len, void *node, void *encoding, int options);

void xmlSetContent(void* node, void *content);

xmlDoc* newEmptyXmlDoc();
xmlElementType getNodeType(xmlNode *node);
char *xmlDocDumpToString(xmlDoc *doc, void *encoding, int format);
char *htmlDocDumpToString(xmlDoc *doc, int format);
void xmlFreeChars(char *buffer);

typedef struct XmlBufferContext {
	char *buffer;
	int buffer_len;
	int data_size;
} XmlBufferContext;

#endif //__CHELPER_H__